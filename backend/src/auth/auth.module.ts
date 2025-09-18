import { Module } from '@nestjs/common';
import { AuthController } from './auth.controller';
import { AuthService } from './auth.service';
import { UtilModule } from 'src/common/util/util.module';
import { JwtModule } from '@nestjs/jwt';
import { JwtStrategy } from './jwt.strategy';
import { MongooseModule } from '@nestjs/mongoose';
import { User, UserSchema } from 'src/common/schema/user.schema';
import { UserService } from 'src/user/user.service';
import { LocalStrategy } from './local.strategy';
import { jwtConstant } from './jwt.constants';

@Module({
  imports: [
    JwtModule.register({
      secret: jwtConstant.secret,
      signOptions: { expiresIn: '1d' },
    }),
    MongooseModule.forFeature([{ name: User.name, schema: UserSchema }]),
    UtilModule,
  ],
  controllers: [AuthController],
  providers: [AuthService, JwtStrategy, UserService, LocalStrategy],
  exports: [JwtStrategy, LocalStrategy, AuthModule],
})
export class AuthModule {}
