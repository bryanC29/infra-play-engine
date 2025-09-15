/* eslint-disable prettier/prettier */
import { Injectable } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import axios from 'axios';
import { Request } from 'express';
import { Result } from './result.dto';

@Injectable()
export class SimulateService {
  constructor(private readonly configService: ConfigService) {}

  async handleSimulation(req: Request) {
    const engineURL =
      this.configService.get<string>('ENGINE_URL') ??
      'http://localhost:3080/simulate';
    const res = await axios.post(engineURL, req.body);
    const data = res.data as Result;

    data.metrics.avgLatency += 'ms';
    data.metrics.avgAvail += '%';

    data.metrics.qps1x.latency += 'ms';
    data.metrics.qps1x.avail += '%';

    data.metrics.qps15x.latency += 'ms';
    data.metrics.qps15x.avail += '%';

    data.metrics.qps2x.latency += 'ms';
    data.metrics.qps2x.avail += '%';

    return data;
  }
}
