#include <stdio.h>
#include <string.h>

#include "hello.h"

static const char *kGreeting = "Hello, ";

enum {
    kOverhead=2
};

struct Hello_T {
    const char *name;
};

bool Hello_create(const char *name, Hello_T *hello) {
    hello->name = name;
    return true;
}

bool Hello_sayHello(const Hello_T *hello, char *out, size_t out_len) {
    size_t len = strlen(hello->name) + strlen(kGreeting) + kOverhead;
    if (out_len < len) {
        return false;
    }
    snprintf(out, out_len, "%s%s!", kGreeting, hello->name);
    return true;
}

const char *Hello_getName(const Hello_T *hello) {
    return hello->name;
}
