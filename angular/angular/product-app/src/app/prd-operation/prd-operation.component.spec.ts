import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PrdOperationComponent } from './prd-operation.component';

describe('PrdOperationComponent', () => {
  let component: PrdOperationComponent;
  let fixture: ComponentFixture<PrdOperationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PrdOperationComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PrdOperationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
