package com.shuaijun.brs.ui.material;

import java.util.List;

public class Material {

    public int code;
    public String msg;
    public List<MaterialDTO> materials;
    public MaterialDTO material;

    public static class MaterialDTO {
        public String mid;
        public String cid;
        public String name;
        public String icon;
    }
}
