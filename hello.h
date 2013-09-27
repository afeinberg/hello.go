#ifndef HELLO_H_
#define HELLO_H_

#include <stdbool.h>

struct Hello_T;

typedef struct Hello_T Hello_T;

bool Hello_create(const char *name, Hello_T *hello);
bool Hello_sayHello(const Hello_T *hello, char *out, size_t out_len);
const char *Hello_getName(const Hello_T *hello);
                          
#endif /* HELLO_H_ */
