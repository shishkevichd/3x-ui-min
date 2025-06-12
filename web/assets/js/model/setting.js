class AllSetting {

    constructor(data) {
        this.webListen = "";
        this.webDomain = "";
        this.webPort = 2053;
        this.webCertFile = "";
        this.webKeyFile = "";
        this.webBasePath = "/";
        this.sessionMaxAge = 60;
        this.pageSize = 50;
        this.expireDiff = 0;
        this.trafficDiff = 0;
        this.remarkModel = "-ieo";
        this.twoFactorEnable = false;
        this.twoFactorToken = "";
        this.xrayTemplateConfig = "";
        this.subEnable = false;
        this.subTitle = "";
        this.subListen = "";
        this.subPort = 2096;
        this.subPath = "/sub/";
        this.subJsonPath = "/json/";
        this.subDomain = "";
        this.externalTrafficInformEnable = false;
        this.externalTrafficInformURI = "";
        this.subCertFile = "";
        this.subKeyFile = "";
        this.subUpdates = 12;
        this.subEncrypt = true;
        this.subShowInfo = true;
        this.subURI = "";
        this.subJsonURI = "";
        this.subJsonFragment = "";
        this.subJsonNoises = "";
        this.subJsonMux = "";
        this.subJsonRules = "";

        if (data == null) {
            return
        }
        ObjectUtil.cloneProps(this, data);
    }

    equals(other) {
        return ObjectUtil.equals(this, other);
    }
}