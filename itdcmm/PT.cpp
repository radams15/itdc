//
// Created by rhys on 05/04/2022.
//

#include <iostream>

#include "include/PT.h"

PT::PT(std::string addr) {
    ptr = pt_new_with_addr((char*) addr.c_str());
}

int PT::heart_rate(){
    return pt_get_heart_rate(ptr);
}

std::string PT::version(){
    return std::string(pt_get_vers(ptr));
}

std::string PT::mac(){
    return std::string(pt_get_mac_addr(ptr));
}

int PT::step_count(){
    return pt_get_step_count(ptr);
}

Motion_t* PT::motion(){
    return pt_get_motion(ptr);
}

void PT::notify(std::string title, std::string subtitle){
    pt_notify(ptr, (char*) title.c_str(), (char*) subtitle.c_str());
}

void PT::set_time(){
    pt_set_time(ptr);
}

void PT::set_time_cust(unsigned long secs){
    pt_set_time_cust(ptr, secs);
}

void PT::set_weather(){
    pt_set_weather(ptr);
}

std::vector<Node_t*> PT::fs_ls(std::string path){
    Directory_t* dir = pt_fs_ls(ptr, (char*) path.c_str());

    return std::vector<Node_t*>(dir->files, dir->files+dir->length);
}

void PT::fs_pull(std::string to_get, std::string save_file){
    pt_fs_pull(ptr, (char*) to_get.c_str(), (char*) save_file.c_str());
}

std::string PT::fs_read(std::string to_get){
    return std::string((char*) pt_fs_read(ptr, (char*) to_get.c_str()));
}

PT::~PT() {
    pt_free(ptr);
}

PT::operator GoUintptr const() {
    return ptr;
}
