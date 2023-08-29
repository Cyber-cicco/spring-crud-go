import { TaskGroup } from './task-group'
import { User } from './user'

export interface UserGroup {
  id: number;
  title: string;
  taskGroupList: TaskGroup[];
  userList: User[];

}
