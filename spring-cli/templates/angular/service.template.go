package angular

var SERVICE_TEMPLATE = 
`import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http
import { environment } from 'src/environments/environment';
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
{%method%}{%by%}({%required_args%}){{%url_changer%}
    return this.http.{%method%}<{%return_type%}>({%url_changed%}{%request_params%}{%body%})
  }
`
var PARAMETER_TEMPLATE =
`{%name%} : {%type%}`

var URL_CHANGER = 
`
    let newURL = this.{%url%};
    newURL = newURL.replace('{{%match%}}', {%match%});`

var URL_DECLARATION = 
`  private {%url_var%} = environement.baseUrl + "{%path%}";`
