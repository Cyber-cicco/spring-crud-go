import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { User } from '../models/user'


@Injectable({
  providedIn: 'root'
})
export class UserHttpService {

  private URL_ = environment.baseUrl + "/user";

  constructor(private http:HttpClient) {}


  get(){
    return this.http.get<User[]>(this.URL_)
  }


  post(dto : User){
    return this.http.post<User>(this.URL_, dto)
  }


  put(dto : User){
    return this.http.put<User>(this.URL_, dto)
  }


  delete(dto : User){
    return this.http.delete<User>(this.URL_, dto)
  }

}
