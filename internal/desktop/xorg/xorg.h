#pragma once

#include <X11/Xlib.h>
#include <X11/XKBlib.h>
#include <X11/Xutil.h>
#include <X11/extensions/Xrandr.h>
#include <X11/extensions/XTest.h>
#include <X11/extensions/Xfixes.h>
#include <stdlib.h>

extern void goCreateScreenSize(int index, int width, int height, int mwidth, int mheight);
extern void goSetScreenRates(int index, int rate_index, short rate);

Display *getXDisplay(void);
int XDisplayOpen(char *input);
void XDisplayClose(void);

void XMove(int x, int y);
void XCursorPosition(int *x, int *y);
void XScroll(int x, int y);
void XButton(unsigned int button, int down);
void XKey(unsigned long key, int down);

void XGetScreenConfigurations();
void XSetScreenConfiguration(int index, short rate);
int XGetScreenSize();
short XGetScreenRate();

void XSetKeyboardModifier(int mod, int on);
char XGetKeyboardModifiers();
XFixesCursorImage *XGetCursorImage(void);

char *XGetScreenshot(int *w, int *h);
