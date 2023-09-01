import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Commentary } from '../models/commentary'


@Injectable({
  providedIn: 'root'
})
export class CommentaryHttpService {

  private URL_ = environment.baseUrl + "/commentary";
  private URL_BLOB = environment.baseUrl + "/commentary/{blob}";

  constructor(private http:HttpClient) {}


  get(){
    return this.http.get<Commentary[]>(this.URL_)
  }


  post(dto : Commentary){
    return this.http.post<Commentary>(this.URL_, dto)
  }


  put(dto : Commentary){
    return this.http.put<Commentary>(this.URL_, dto)
  }


  delete(dto : Commentary){
    return this.http.delete<Commentary>(this.URL_, dto)
  }


  getBlob(blob : number){
    let newURL = this.URL_BLOB;
    newURL = newURL.replace('{blob}', blob.toString());
    return this.http.get<Commentary>(newURL)
  }

}
