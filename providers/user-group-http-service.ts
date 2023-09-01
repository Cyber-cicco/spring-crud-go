import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { UserGroup } from '../models/user-group'


@Injectable({
  providedIn: 'root'
})
export class UserGroupHttpService {

  private URL_ = environment.baseUrl + "/userGroup";

  constructor(private http:HttpClient) {}


  get(){
    return this.http.get<UserGroup[]>(this.URL_)
  }


  post(dto : UserGroup){
    return this.http.post<UserGroup>(this.URL_, dto)
  }


  put(dto : UserGroup){
    return this.http.put<UserGroup>(this.URL_, dto)
  }


  delete(dto : UserGroup){
    return this.http.delete<UserGroup>(this.URL_, dto)
  }

}
