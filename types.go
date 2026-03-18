package main

//import "time"
///
type HlmSpec struct {
			Env string `yaml:"env,omitempty"`
			StsSingle bool `yaml:"stsSingle,omitempty"`
			Application string `yaml:"application,omitempty"`
			Cloudns string `yaml:"cloudns,omitempty"`
			Helmv3tpl *map[string]map[string]Releasees `yaml:"Helmv3tpl,omitempty"`
			Coredns string `yaml:"coredns,omitempty"`
			Global /**/ `yaml:"global,omitempty"`
}
type Releasees struct {
			Rle *[]map[string]any `yaml:"_role,omitempty"`
			RoleBindingpls *[]map[string]any `yaml:"_roleBinding+,omitempty"`
			_ *[]map[string]any `yaml:"_,omitempty"`
			Cjpls *[]map[string]any `yaml:"cj+,omitempty"`
			ClusterRole *[]map[string]any `yaml:"clusterRole,omitempty"`
			ClusterroleBindingpls *[]map[string]any `yaml:"clusterroleBinding+,omitempty"`
	Cmpls *[]struct {
			Files string `yaml:"files,omitempty"`
	} `yaml:"cm+,omitempty"`
			Crd *[]map[string]any `yaml:"crd,omitempty"`
			Deploypls *[]map[string]any `yaml:"deploy+,omitempty"`
			Drpls *[]map[string]any `yaml:"dr+,omitempty"`
	Gwpls *[]struct {
			Hosts string `yaml:"hosts,omitempty"`
		Vps *[]struct {
			Ptc string `yaml:"ptc,omitempty"`
			Number string `yaml:"number,omitempty"`
			Name string `yaml:"name,omitempty"`
		} `yaml:"vps,omitempty"`
			CredentialName string `yaml:"credentialName,omitempty"`
	} `yaml:"gw+,omitempty"`
	Hpa *[]struct {
			Max int32 `yaml:"max,omitempty"`
	} `yaml:"hpa,omitempty"`
			Nspls *[]map[string]any `yaml:"ns+,omitempty"`
	Pvpls *[]struct {
			StgSize string `yaml:"stgSize,omitempty"`
			AccMode string `yaml:"accMode,omitempty"`
			VolHandle string `yaml:"volHandle,omitempty"`
			Servers string `yaml:"servers,omitempty"`
			Paths string `yaml:"paths,omitempty"`
			Drivers string `yaml:"drivers,omitempty"`
			MntOpts *[]string `yaml:"mntOpts,omitempty"`
	} `yaml:"pv+,omitempty"`
	Pvc *[]struct {
			AccMode string `yaml:"accMode,omitempty"`
			VolumeName string `yaml:"volumeName,omitempty"`
			StgSize string `yaml:"stgSize,omitempty"`
	} `yaml:"pvc,omitempty"`
	Sapls *[]struct {
			Annotaions map[string]string `yaml:"annotations,omitempty"`
	} `yaml:"sa+,omitempty"`
	Sepls *[]struct {
			Ip string `yaml:"ip,omitempty"`
		Vps *[]struct {
			Ptc string `yaml:"ptc,omitempty"`
			Number string `yaml:"number,omitempty"`
			Name string `yaml:"name,omitempty"`
		} `yaml:"vps,omitempty"`
			Is bool `yaml:"is,omitempty"`
	} `yaml:"se+,omitempty"`
	Secretpls *[]struct {
		Data struct {
			Crt string `yaml:"tls.crt,omitempty"`
			Cacerts	 string `yaml:"cacerts.pem,omitempty"`
			Key string `yaml:"tls.key,omitempty"`
		} `yaml:"data,omitempty"`
			Files string `yaml:"files,omitempty"`
			// * `yaml:"*,omitempty"`
	} `yaml:"secret+,omitempty"`
	StorageClasspls *[]struct {
		Annotaions struct {
			DefaultsStg bool `yaml:"storageclass.kubernetes.io/is-default-class,omitempty"` //onlyOne
		} `yaml:"annotations,omitempty"`
	} `yaml:"storageClass+,omitempty"`
	Svc *[]struct {
			ClusterIP string `yaml:"clusterIP,omitempty"`
		Vps *[]struct {
			Ptc string `yaml:"ptc,omitempty"`
			S string `yaml:"s,omitempty"`
			Number string `yaml:"number,omitempty"`
			T string `yaml:"t,omitempty"`
			Name string `yaml:"name,omitempty"`
		} `yaml:"vps,omitempty"`
			ExternalIPs *[]map[string]any `yaml:"externalIPs,omitempty"`
	} `yaml:"svc,omitempty"`
	Vspls *[]struct {
		Vps *[]struct {
			Ptc string `yaml:"ptc,omitempty"`
			S string `yaml:"s,omitempty"`
			Number string `yaml:"number,omitempty"`
			T string `yaml:"t,omitempty"`
			Name string `yaml:"name,omitempty"`
		} `yaml:"vps,omitempty"`
	} `yaml:"vs+,omitempty"`
}
type Global struct {
			Reg string `yaml:"reg,omitempty"`
			Label *map[string]string `yaml:"label,omitempty"`
			Tag string `yaml:"tag,omitempty"`
}

func init() {
	//
}