import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Tag } from '../models/tag'


@Injectable({
  providedIn: 'root'
})
export class TagHttpService {

  private URL_ = environment.baseUrl + "/tag";

  constructor(private http:HttpClient) {}


  get(){
    return this.http.get<Tag[]>(this.URL_)
  }


  post(dto : Tag){
    return this.http.post<Tag>(this.URL_, dto)
  }


  put(dto : Tag){
    return this.http.put<Tag>(this.URL_, dto)
  }


  delete(dto : Tag){
    return this.http.delete<Tag>(this.URL_, dto)
  }

}
