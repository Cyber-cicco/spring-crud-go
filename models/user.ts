import { UserGroup } from './user-group'
import { Commentary } from './commentary'
import { Task } from './task'

export interface User {
  id: number;
  nom: string;
  prenom: string;
  email: string;
  password: string;
  userGroupList: UserGroup[];
  commentaryList: Commentary[];
  taskList: Task[];

}
