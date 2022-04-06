//
// Created by rhys on 05/04/2022.
//

#ifndef ITDC_PT_H
#define ITDC_PT_H

#include <libitdc.h>

#ifdef __cplusplus

#include <string>
#include <vector>

class PT{
private:

    GoUintptr ptr;

public:
    PT(std::string addr="/tmp/itd/socket");

    int heart_rate();
    std::string version();
    std::string mac();
    int step_count();
    Motion_t* motion();
    void notify(std::string title, std::string subtitle);
    void set_time();
    void set_time_cust(unsigned long secs);
    void set_weather();

    std::vector<Node_t*> fs_ls(std::string path);
    void fs_pull(std::string to_get, std::string save_file);
    std::string fs_read(std::string to_get);

    explicit operator GoUintptr const();

    ~PT();
};

#endif

#endif //ITDC_PT_H
