import { TaskGroup } from './task-group'
import { User } from './user'
import { Tag } from './tag'
import { Commentary } from './commentary'

export interface Task {
  id: number;
  title: string;
  description: string;
  status: string;
  taskGroup: TaskGroup;
  dateCreation: Date;
  author: User;
  dateEcheance: Date;
  tagList: Tag[];
  next: Task;
  commentaryList: Commentary[];

}
