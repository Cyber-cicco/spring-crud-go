import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { TaskGroup } from '../models/task-group'


@Injectable({
  providedIn: 'root'
})
export class TaskGroupHttpService {

  private URL_ = environment.baseUrl + "/taskGroup";

  constructor(private http:HttpClient) {}


  get(){
    return this.http.get<TaskGroup[]>(this.URL_)
  }


  post(dto : TaskGroup){
    return this.http.post<TaskGroup>(this.URL_, dto)
  }


  put(dto : TaskGroup){
    return this.http.put<TaskGroup>(this.URL_, dto)
  }


  delete(dto : TaskGroup){
    return this.http.delete<TaskGroup>(this.URL_, dto)
  }

}
