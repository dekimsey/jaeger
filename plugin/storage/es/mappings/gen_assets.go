// Code generated by "esc -pkg mappings -o plugin/storage/es/mappings/gen_assets.go -ignore assets -prefix plugin/storage/es/mappings plugin/storage/es/mappings"; DO NOT EDIT.

package mappings

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/.nocover": {
		name:    ".nocover",
		local:   "plugin/storage/es/mappings/.nocover",
		size:    43,
		modtime: 1595973361,
		compressed: `
H4sIAAAAAAAC/youSSzJzFYoSEzOTkxPVcjILy4pVkgsLcnXTU/NSy1KLElNUUjLzEkt1uMCBAAA//8y
IKK1KwAAAA==
`,
	},

	"/jaeger-dependencies-7.json": {
		name:    "jaeger-dependencies-7.json",
		local:   "plugin/storage/es/mappings/jaeger-dependencies-7.json",
		size:    283,
		modtime: 1595973361,
		compressed: `
H4sIAAAAAAAC/2zPz0vDQBDF8Xv+imXxVNrFi5fcqlYU/EWK52GbfU1HknHdmYBQ8r9LRA/S3t/nC+9Y
OedZEr4oRzMUUV87v3iP6FBWCRmSIC1DVwu/nNcKM5ZOfT3jPx5kHHYo9LEnPcSS5szFkej57el609DL
HW3v183tlmhanmcFuec2nsJm8/r4cLM+oUPMmaULAjUk2jP6pKHngc3XV5f/tgWfI9Q0tLE9IEDiroev
rYyonPvp+t/efGyqpuo7AAD//66cHf8bAQAA
`,
	},

	"/jaeger-dependencies.json": {
		name:    "jaeger-dependencies.json",
		local:   "plugin/storage/es/mappings/jaeger-dependencies.json",
		size:    277,
		modtime: 1595973361,
		compressed: `
H4sIAAAAAAAC/2zPzUoDMRTF8f08Rbi4Km1w4ya7qhUFv5ji+pJOTqeRTIy5d0Ao8+4y4kbG/fn94Zwb
Y0gxlOQV5Ayt3j161E1AQQ7IXYRsVrSedwLVmHshNzNjKOaAL5vH4YDKH0eWk69ByJmLM/Pz29P1ruWX
O97fb9vbPfO0/p9VlBQ7v4Tt7vXx4Wa7oIMvJebeZogi8DEiBbEpDlHJXV3+2VZ8jhAV2/nuBIvsDwnk
tI5ojPnp0m9vPjY1U/MdAAD//5ZQx/QVAQAA
`,
	},

	"/jaeger-service-7.json": {
		name:    "jaeger-service-7.json",
		local:   "plugin/storage/es/mappings/jaeger-service-7.json",
		size:    878,
		modtime: 1595973361,
		compressed: `
H4sIAAAAAAAC/8ySwW7aQBCG736K1agnBFZViR72RluqVmppBcopikaDPdibeNeb3YEEIb97ZGTAhHDL
IReP5H++35/t3SVKgXE5P6MnEQ4uglYwuCcuOIwih43JeDSAYbsYWcS4IoJuuQOZurVdcsB6hbGkkLcN
n3aIs5u/36Zz/PcTF78m8x8LxGb4NhbYVyajS3A+/f/n9/fJBWrJe+OK1HEUznFluMpjWhlrBPT489lu
4Mc1R4lpRlnJKTtaVgxawpoTpfa90PWdXizfOrImQ2HrKxKOoG/3iVK7brbfw5NDoSKiJX9gu6yrPL+r
FMjWM2h44O1THXIYnqemcHVgpGW9YdBfxl97cdPfBU9SoiXJStAgVKQDOMZN8oroOftQZxzjh9DuXNJr
+vt51/1NH2rPQQzHkxx0B3RGlvvK13Wvqh41oX0Miande7Qmh2uTNMlLAAAA//8YcMrbbgMAAA==
`,
	},

	"/jaeger-service.json": {
		name:    "jaeger-service.json",
		local:   "plugin/storage/es/mappings/jaeger-service.json",
		size:    1060,
		modtime: 1595973361,
		compressed: `
H4sIAAAAAAAC/8yTT2/UMBDF7/kU1ojTamshpHLwrUARSFDQVpwQGs3Gs1mD7RjbKayqfHfk4pKkW+2J
Q3PIn/F78/xz7NtGCMjsgqXMoASsvhN3HM8SxxvT8tkK1kWSOGfjuwSqOIQA4zX/ln5wW47Y7zDtKeoE
Sjy7Rbz68vHV5QY/vcXrdxebN9eI4/pxW+RgTUvHxs3l5w/vX18cWR2FYHwnPafMGneGrU7SGmcyqPPn
C23knwOnnGRL7Z4le9paBpXjwEc9OUp98ORMC2pHNnEjxF0y1MQJHTXvaLAZ7yulRtZOn0LA3zA9NStX
RRECahbeL30C9fWfeWpTVj6Qx0xdQkdhHnE3Wif3sF5+6iEwKPjBh1991LB+OG4630dG2vY3DOrF+cuF
YFzqIVDeo6Pc7kFBpk6uYCYYm0d8C4oQ+5ZTemIgdVbyFFB9+9bMukE9HbMNEGIfOGbDabENqvCKHC/R
TmGdQJrhQAmkbHr//7o382e5j83Y/AkAAP//qd2MzCQEAAA=
`,
	},

	"/jaeger-span-7.json": {
		name:    "jaeger-span-7.json",
		local:   "plugin/storage/es/mappings/jaeger-span-7.json",
		size:    3420,
		modtime: 1595973361,
		compressed: `
H4sIAAAAAAAC/+xWXW+UQBR951eQG5+aLTEm9YG3amtsYqtp65Mxk7twYaedL2fuVjcN/93A0hYKbE2k
xhhflix3zuHcmXsO3EZxDNLk9EM4ZCZvAqQx7F0hleT3g0OzvweLuF4WiFmaMkBao+5wiVnrJXlhCxFW
6PMa/+JWiLPPp2+Oz8XHd+Li/eH50YUQ1WIc5skpmeEQeH786cPJ28MBVKNz0pSJocCUi0KSykOipJYM
6cHL3lpP39YUOCQZZitKyOBSEaTs1xTFccMLLd9DY/nGoJaZYNJOIVOA9EtTiePb9lrvh0MjGMsgNLo7
bFtrKft34xh44whSuKbNd+tzWPSrsjTWk8ClvSFIXx287pSr7lpwyCuhkbMVpMBYJntwX66iR4iOZudt
RiH8FbJbLcmU/Ob6tT1N560jz5LCgzhgjxmdHHXlTkudlHkvERx6Mnzh0MxIGualq7cBWVpzhprmE8no
+VKOMyprSpiGnEqlZBgD5sjU1VFYr5EhBXI2Wwm9BQ6Y8/W2w1/XUigsRxVIw3WQDRHKjgO2mdIV3YYB
pAWqQIuelwYjuaWSmgKjdlM+6jYxNMk2z6awA4E7Re4U2hSvaTO8+5Tln7T9oKsGcYNqTX/saYzlZUP7
PM+Lpv5V00l8l3m9yZuco0D+Rmb02OizJjLjZNrb5RVlDLug/6f0n5xSTwV5Mhk9W0R6Ksa6nm+sh18G
s/IPX+q/S7/jOB55dNyfveXPdm4jPpxtT0d9N2fQzT1xE5+s9W8VVdHPAAAA//+SuQbQXA0AAA==
`,
	},

	"/jaeger-span.json": {
		name:    "jaeger-span.json",
		local:   "plugin/storage/es/mappings/jaeger-span.json",
		size:    3830,
		modtime: 1595973361,
		compressed: `
H4sIAAAAAAAC/+xW0W/TPhB+z18RnX5PUxf9hDQe8jbYEJPYQNt4Qsi6JpfUm2Mb+zqopv7vKE1Lm9ZJ
QGoQEvShbWx/391n333xcxTHwFRZhUyQxnDygFSSO/UW9ekJTOp5T8xSlx7Senkcg9Q5fUv0vJqSE6YQ
foYu95DG/z0LcfPx+tXlrXj/Rty9Pb+9uBNiOQnDHFklMzwE3l5+eHf1+vwAWqG1UpeJJs+Ui0KSyn2i
ZCUZ0rP/W2sdfZmTZ59kmM0oIY1TRZCym9MBJ7kkX2isZAZpgcpTFMeryLCOuJUucipwrlhsRuoxVGr7
GMfQBMu3ZPVnLSWOYR1LbPbdQ/rpB3hLU++8RS0YSy8qtLshVrPr5PbH6xNdWIIUHmnx1bgcJvvzstTG
kcCpeSJIX5y9bC1YtteDRZ6JCjmbQQqMZXICOwuWUQDXUmGdycj7P0zIOqukT9D63+doh211KDunb52x
5FiSb9UAO8zo6qKtqU9Pj5YdHWDRkeY7i3oEcj8Obb1ByNLoG6zo+EkzOr6XXczK6BL6gddSKenD8Lw2
xlZWhXEVMqRA1mQzUTXgYIR83uj+1cwKhWVHPlJzbdBhnDJdsMYy20LanjfZa9lAUTeEsiLPWNnudm0L
C3ViY93dDIF0B1IeSHs1/UiL0Piwz/yE1wRUrlBPqOb026Mylvcr8jHjRn3Py6FXxMaA9+q1p/I8uSeZ
0aGJjPCqYOx5HZnpA2UMQwT/6vuvrm9HBTnSGY1uyY6K8G4cuylCd5oR4oSuIccIM3hkB13f1fF7oNFP
ONjZR971jk4+vrGOU6u9N/jmt/5eRsvoewAAAP//W45CgfYOAAA=
`,
	},

	"/": {
		name:  "/",
		local: `plugin/storage/es/mappings`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"plugin/storage/es/mappings": {
		_escData["/.nocover"],
		_escData["/jaeger-dependencies-7.json"],
		_escData["/jaeger-dependencies.json"],
		_escData["/jaeger-service-7.json"],
		_escData["/jaeger-service.json"],
		_escData["/jaeger-span-7.json"],
		_escData["/jaeger-span.json"],
	},
}
