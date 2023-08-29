import { Task } from './task'

export interface Tag {
  id: number;
  title: string;
  color: string;
  taskList: Task[];

}
