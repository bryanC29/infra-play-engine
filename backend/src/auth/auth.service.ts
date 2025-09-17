/* eslint-disable prettier/prettier */
import { HttpStatus, Injectable, Res } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Response } from 'express';
import { Model } from 'mongoose';
import * as bcrypt from 'bcryptjs';
import { User, UserDocument } from 'src/common/schema/user.schema';
import { LoginDTO, RegisterDTO } from './auth.dto';
import { UtilService } from 'src/common/util/util.service';
import { JwtService } from '@nestjs/jwt';
import { UserService } from 'src/user/user.service';

@Injectable()
export class AuthService {
  constructor(
    @InjectModel(User.name) private readonly userModel: Model<UserDocument>,
    private readonly utilService: UtilService,
    private readonly jwtService: JwtService,
    private readonly userService: UserService,
  ) {}

  async validateUser(email: string, password: string): Promise<User | false> {
    const user = await this.userService.findUser(email);

    if (!user) {
      return false;
    }

    const passEqual = bcrypt.compareSync(password, user.password);
    if (!passEqual) {
      return false;
    }

    return user;
  }

  async register(body: RegisterDTO, @Res() res: Response) {
    const oldUser = await this.userService.findUser(body.email);
    if (oldUser != null) {
      this.utilService.sendUserAlreadyRegistered(res);
      return;
    }

    try {
      const hashedPassword = await bcrypt.hash(body.password, 10);
      const newUser = new this.userModel({
        ...body,
        password: hashedPassword,
      });

      await newUser.save();

      const payload = { email: newUser.email, name: newUser.name };
      const token = await this.jwtService.signAsync(payload);

      this.utilService.sendUserRegistered(res, token);

      return;
    } catch {
      this.utilService.sendUnknownError(res);
      return;
    }
  }

  async login(body: LoginDTO, @Res() res: Response) {
    try {
      const user = await this.validateUser(body.email, body.password);

      if (!user) {
        this.utilService.sendUserNotFound(res);
        return;
      }

      const payload = { name: user.name, email: user.email };
      const token = await this.jwtService.signAsync(payload);

      this.utilService.sendUserLogged(res, token);

      return;
    } catch {
      this.utilService.sendUnknownError(res);
    }

    return;
  }

  logout(@Res() res: Response) {
    res.clearCookie('token');
    res.setHeader('Authorization', 'Bearer null');
    res.status(200).send({
      error: '',
      message: ['user logout seccessfully'],
      statusCode: HttpStatus.OK,
    });
  }
}
