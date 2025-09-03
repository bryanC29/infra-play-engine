/* eslint-disable prettier/prettier */
import { Injectable } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import axios from 'axios';
import { Request } from 'express';

@Injectable()
export class SimulateService {
  constructor(private readonly configService: ConfigService) {}

  async handleSimulation(req: Request) {
    const engineURL =
      this.configService.get<string>('ENGINE_URL') ??
      'http://localhost:3080/simulate';
    const res = await axios.post(engineURL, req.body);

    return res.data;
  }
}
