import { User } from './user'

export interface Commentary {
  id: number;
  content: string;
  author: User;
  rating: string;
  responseList: Commentary[];

}
