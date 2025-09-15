/* eslint-disable prettier/prettier */
import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';

@Schema({ _id: false })
export class Submission {
  @Prop({ required: true })
  submissionId: string;

  @Prop({ required: true })
  avgSuccess: number;

  @Prop({ required: true })
  avgFailed: number;

  @Prop({ required: true })
  avgLatency: string;

  @Prop({ required: true })
  avgAvail: string;
}

@Schema()
export class User {
  @Prop({ required: true })
  name: string;

  @Prop({ unique: true, required: true })
  email: string;

  @Prop({ required: true })
  password: string;

  @Prop({ type: [Submission], default: [] })
  submission: Submission[];
}

export const UserSchema = SchemaFactory.createForClass(User);
export const SubmissionSchema = SchemaFactory.createForClass(Submission);
export type UserDocument = User & Document;
