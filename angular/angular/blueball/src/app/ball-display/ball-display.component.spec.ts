import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BallDisplayComponent } from './ball-display.component';

describe('BallDisplayComponent', () => {
  let component: BallDisplayComponent;
  let fixture: ComponentFixture<BallDisplayComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BallDisplayComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(BallDisplayComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
