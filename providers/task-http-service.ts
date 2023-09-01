import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Task } from '../models/task'


@Injectable({
  providedIn: 'root'
})
export class TaskHttpService {

  private URL_ = environment.baseUrl + "/task";

  constructor(private http:HttpClient) {}


  get(){
    return this.http.get<Task[]>(this.URL_)
  }


  post(dto : Task){
    return this.http.post<Task>(this.URL_, dto)
  }


  put(dto : Task){
    return this.http.put<Task>(this.URL_, dto)
  }


  delete(dto : Task){
    return this.http.delete<Task>(this.URL_, dto)
  }

}
