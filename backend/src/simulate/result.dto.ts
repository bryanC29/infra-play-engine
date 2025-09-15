/* eslint-disable prettier/prettier */
export type Result = {
  submissionId: string;
  baseQps: number;
  totalNodes: number;
  metrics: Metrics;
};

type Metrics = {
  qps1x: Qps;
  qps15x: Qps;
  qps2x: Qps;
  avgSuccess: number;
  avgFailed: number;
  avgLatency: string;
  avgAvail: string;
};

type Qps = {
  success: number;
  failed: number;
  latency: string;
  avail: string;
};
