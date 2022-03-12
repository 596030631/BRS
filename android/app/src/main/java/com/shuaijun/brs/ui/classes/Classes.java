package com.shuaijun.brs.ui.classes;

import androidx.annotation.NonNull;

import java.util.List;

public class Classes {

    public int code;
    public String msg;
    public List<ClassesDTO> classes;

    public static class ClassesDTO {
        public String cid;
        public String name;
        public String pid;

        @NonNull
        @Override
        public String toString() {
            return "ClassesDTO{" +
                    "cid='" + cid + '\'' +
                    ", name='" + name + '\'' +
                    ", pid='" + pid + '\'' +
                    '}';
        }
    }

    @NonNull
    @Override
    public String toString() {
        return "Classes{" +
                "code=" + code +
                ", msg='" + msg + '\'' +
                ", classes=" + classes +
                '}';
    }
}
