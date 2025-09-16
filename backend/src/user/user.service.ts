/* eslint-disable prettier/prettier */
import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { User, UserDocument } from 'src/common/schema/user.schema';

@Injectable()
export class UserService {
  constructor(
    @InjectModel(User.name) private readonly userModel: Model<UserDocument>,
  ) {}

  async findUser(email: string): Promise<User | null> {
    try {
      const user = await this.userModel.findOne({ email: email });

      if (!user) {
        return null;
      }

      return user;
    } catch {
      return null;
    }
  }
}
