#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <memory.h>

int main (int argc, char* argv[]) {

    char currentPath[256];
    char path[256];

    // getcwd is the equivalent of pwd command in bash
    getcwd(currentPath, sizeof(currentPath));
    printf("Your current path is : %s\n",currentPath);

    printf("Path : \n");
    fgets(path, sizeof(path), stdin);

    // we have to remove the carriage return to perform a bash command
    char cleanPath[256];
    strncpy(cleanPath, path, strlen(path)-1);

    // snprintf is used in order to create the bash command with an argument
    char buffer[256];
    snprintf(buffer, sizeof(buffer), "find %s -type f | wc -l", cleanPath);

    // the command is launched and the output is stored in a variable thanks to fscanf
    int count;
    FILE *output = popen(buffer, "r");
    fscanf(output, "%d", &count);
    pclose(output);

    printf("Files : %d\n", count);
    return 0;
}