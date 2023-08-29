import { UserGroup } from './user-group'
import { Task } from './task'

export interface TaskGroup {
  id: number;
  title: string;
  description: string;
  userGroupList: UserGroup[];
  taskList: Task[];

}
