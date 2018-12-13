import { Component, OnInit, Input, OnChanges, SimpleChanges, Output, EventEmitter } from '@angular/core';
import { DataService } from '../../services/data.service';
import { ContivNodeDataModel } from '../../models/contiv-node-data-model';
import { VppInterfaceModel } from '../../models/vpp/vpp-interface-model';
import { Subscription } from 'rxjs';
import { VppBdModel } from '../../models/vpp/vpp-bd-model';
import { TopologyHighlightService } from '../../../d3-topology/topology-viz/topology-highlight.service';
import { Router, ActivatedRoute } from '@angular/router';

interface BdRow {
  node: string;
  bviName: string;
  bviIp: string;
  vrf: number;
  vni: number;
  podsCount: number;
  vxlansCount: number;
}

interface PodRow {
  name: string;
  iface: string;
  ip: string;
  node: string;
}

interface VxlanRow {
  srcNode: string;
  dstNode: string;
  name: string;
  srcIP: string;
  dstIP: string;
  srcBvi: string;
  dstBvi: string;
}

@Component({
  selector: 'app-bridge-domain-control',
  templateUrl: './bridge-domain-control.component.html',
  styleUrls: ['./bridge-domain-control.component.css']
})
export class BridgeDomainControlComponent implements OnInit, OnChanges {

  @Input() tableType: string;
  @Output() detailShowed: EventEmitter<{
    nodeId: string, dstId?: string, type: 'vppPod' | 'bvi' | 'vxtunnel'
  }> = new EventEmitter<{nodeId: string, dstId?: string, type: 'vppPod' | 'bvi' | 'vxtunnel'}>();

  public domains: ContivNodeDataModel[];
  public summaryObj: BdRow[];
  public podsObj: PodRow[];
  public tunnelsObj: VxlanRow[];
  public vxlans: VppInterfaceModel[];
  public bd: VppBdModel;

  public isSummary: boolean;
  public isPods: boolean;
  public isTunnels: boolean;

  private subscriptions: Subscription[];

  constructor(
    private router: Router,
    private route: ActivatedRoute,
    private dataService: DataService,
    private topologyHighlightService: TopologyHighlightService
  ) { }

  ngOnInit() {
    this.subscriptions = [];
    this.summaryObj = [];
    this.podsObj = [];
    this.tunnelsObj = [];

    this.isSummary = false;
    this.isPods = false;
    this.isTunnels = false;

    this.subscriptions.push(
      this.dataService.isContivDataLoaded.subscribe(isLoaded => {
        if (isLoaded) {
          this.bd = this.dataService.contivData.contivData[0].bd[0];
          this.domains = this.dataService.contivData.contivData;
          this.summaryObj = this.domains.map(d => {
            const bvi = d.getBVI();
            const vxlans = d.getVxlans();
            const row: BdRow = {
              node: d.node.name,
              bviName: d.vswitch.name + '-bvi',
              bviIp: bvi.IPS,
              vrf: bvi.vrf,
              vni: vxlans[0].vni,
              podsCount: d.getTapInterfaces().length,
              vxlansCount: vxlans.length
            };

            return row;
          });

          this.domains.forEach(d => {
            d.vppPods.forEach(pod => {
              if (!pod.name.includes('coredns')) {
                const row: PodRow = {
                  name: pod.name,
                  iface: pod.tapInternalInterface,
                  ip: pod.podIp,
                  node: pod.node
                };

                this.podsObj.push(row);
              }
            });
          });

          this.domains.forEach(d => {
            d.getVxlans().forEach(vx => {
              const srcNode = this.dataService.contivData.getNodeByIpamIp(vx.srcIP);
              const dstNode = this.dataService.contivData.getNodeByIpamIp(vx.dstIP);

              const srcDomain = this.dataService.contivData.getDomainByNodeId(srcNode.name);
              const dstDomain = this.dataService.contivData.getDomainByNodeId(dstNode.name);

              const row: VxlanRow = {
                srcIP: vx.srcIP,
                srcNode: srcNode.name,
                dstIP: vx.dstIP,
                dstNode: dstNode.name,
                name: vx.name,
                srcBvi: srcDomain.vswitch.name + '-bvi',
                dstBvi: dstDomain.vswitch.name + '-bvi'
              };

              this.tunnelsObj.push(row);
            });
          });
        } else {
          this.summaryObj = [];
          this.podsObj = [];
          this.tunnelsObj = [];
        }
      })
    );
  }

  public showDetail(nodeId: string, type: 'vppPod' | 'bvi' | 'vxtunnel', dstId?: string) {
    this.detailShowed.emit({nodeId: nodeId, dstId: dstId, type: type});
  }

  public highlightNode(nodeId: string) {
    this.topologyHighlightService.highlightNode(nodeId);
  }

  public highlightBvi(bviId: string) {
    this.topologyHighlightService.highlightBVI(bviId);
  }

  public clearHighlight() {
    this.topologyHighlightService.clearSelections();
  }

  public highlightTunnel(srcId: string, dstId: string) {
    this.topologyHighlightService.highlightLinkBetweenNodes(srcId, dstId);
  }

  ngOnChanges(changes: SimpleChanges) {
    if (!changes.tableType.firstChange) {
      switch (this.tableType) {
        case 'summary':
          this.isSummary = true;
          this.isPods = false;
          this.isTunnels = false;
          break;
        case 'pods':
          this.isSummary = false;
          this.isPods = true;
          this.isTunnels = false;
          break;
        case 'tunnels':
          this.isSummary = false;
          this.isPods = false;
          this.isTunnels = true;
          break;
      }
    }
  }
}
