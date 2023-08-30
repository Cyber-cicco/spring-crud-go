package angular

var SERVICE_TEMPLATE = 
`import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http
{%imports%}

@Injectable({
  providedIn: 'root'
})
export class {%class_name%} {

{%urls%}

  constructor(private http:HttpClient) {}

{%http%}
}
';`

var SERVICE_METHOD_TEMPLATE =
`
{%method%}{%target_name%}{%by%}({%required_args%}){{%url_changer%}
    return this.http.{%method%}<{%return_type%}>({%url_changed%}{%request_params%}{%body%})
  }
`
