/* eslint-disable prettier/prettier */
import { HttpStatus, Injectable } from '@nestjs/common';
import { ResponseDTO } from './response.dto';
import { Response } from 'express';

@Injectable()
export class UtilService {
  private response: ResponseDTO;
  constructor() {
    this.response = {
      message: [''],
      error: '',
      statusCode: 0,
    };
  }

  private setCookie(res: Response, token: string) {
    res.setHeader('Authorization', `Bearer ${token}`);
    res.cookie('token', token, {
      httpOnly: true,
      maxAge: 1000 * 60 * 60 * 24 * 7,
    });
  }

  sendUserLogged(res: Response, token: string) {
    const code = HttpStatus.OK;

    this.response.error = '';
    this.response.message[0] = 'user logged in successfully';
    this.response.statusCode = code;

    this.setCookie(res, token);

    res.status(code).send(this.response);
  }

  sendUserRegistered(res: Response, token: string): void {
    const code = HttpStatus.CREATED;

    this.response.error = '';
    this.response.message[0] = 'user registered successfully';
    this.response.statusCode = code;

    this.setCookie(res, token);

    res.status(code).send(this.response);
  }

  sendUserAlreadyRegistered(res: Response) {
    const code = HttpStatus.CONFLICT;

    this.response.error = 'error registering user';
    this.response.message[0] = 'provided email already registered';
    this.response.statusCode = code;

    res.status(code).send(this.response);
  }

  sendUserNotFound(res: Response): void {
    const code = HttpStatus.NOT_FOUND;

    this.response.error = 'user not found';
    this.response.message[0] = 'error occured while fetching user information';
    this.response.statusCode = code;

    res.status(code).send(this.response);
  }

  sendUnauthorized(res: Response): void {
    const code = HttpStatus.UNAUTHORIZED;

    this.response.error = 'unauthorized';
    this.response.message[0] = 'action requires appropriate authorization';
    this.response.message[1] =
      'please login using authorized account to perform action';
    this.response.statusCode = code;

    res.status(code).send(this.response);
  }

  sendUnknownError(res: Response): void {
    const code = HttpStatus.INTERNAL_SERVER_ERROR;

    this.response.error = 'unknown error';
    this.response.message[0] = 'an unknown error occured';
    this.response.message[1] = 'please try again later';
    this.response.statusCode = code;

    res.status(code).send(this.response);
  }
}
