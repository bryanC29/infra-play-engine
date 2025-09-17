/* eslint-disable prettier/prettier */
import {
  Body,
  Controller,
  Get,
  HttpCode,
  HttpStatus,
  Post,
  Res,
  UseGuards,
} from '@nestjs/common';
import { AuthService } from './auth.service';
import { Response } from 'express';
import { LoginDTO, RegisterDTO } from './auth.dto';
import { LocalAuthGuard } from './auth.guard';

@Controller('auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @HttpCode(HttpStatus.CREATED)
  @Post('register')
  async registerUser(@Body() body: RegisterDTO, @Res() res: Response) {
    return await this.authService.register(body, res);
  }

  @UseGuards(LocalAuthGuard)
  @HttpCode(HttpStatus.OK)
  @Post('login')
  async loginUser(@Body() body: LoginDTO, @Res() res: Response) {
    return await this.authService.login(body, res);
  }

  @HttpCode(HttpStatus.OK)
  @Get('logout')
  logoutUser(@Res() res: Response) {
    return this.authService.logout(res);
  }
}
