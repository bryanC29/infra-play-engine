/* eslint-disable prettier/prettier */
import { Controller, HttpCode, Post, Req } from '@nestjs/common';
import { SimulateService } from './simulate.service';
import { Request } from 'express';

@Controller('simulate')
export class SimulateController {
  constructor(private readonly simService: SimulateService) {}

  @Post()
  @HttpCode(200)
  async simulation(@Req() req: Request) {
    return await this.simService.handleSimulation(req);
  }
}
