import { Directive,Input,TemplateRef,ViewContainerRef} from '@angular/core';

@Directive({
  selector: '[appStructural]'
})
export class StructuralDirective {

  constructor( private template: TemplateRef<any>,
    private view: ViewContainerRef) {
    
   }
  @Input() set appStructural(condition: boolean) {
    if (!condition ) {
      this.view.createEmbeddedView(this.template)
    } else if (condition ) {
      this.view.clear();
    }
  }

}
