import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-output-format',
  templateUrl: './output-format.component.html',
  styleUrls: ['./output-format.component.css']
})
export class OutputFormatComponent implements OnInit {
  data:string="";
  fontSize!:string;
  colors:string[]=["yellow","violet","blue"];
  BackColor!:string;
  TextFontColor!:string;
  constructor() { }
  ngOnInit(): void {
  }

}
