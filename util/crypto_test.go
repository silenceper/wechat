package util

import (
	"encoding/base64"
	"testing"
)

func TestECBDecrypt(t *testing.T) {
	key := []byte("04dc8389d85cb3093781ff6c6f994cd9")
	data := "1qXIZ5HeCS53vq6vekntsRBkPyaqd+c+282L+yNmCU6qBiLqnJyDFdwAeeM385TJ4KDD3p4EMJMdRnuU68I0tr2nO1h3NypEE0g+wxM4NKplol64MUCCX0En8OA1vRonQx1I6jV3OIXOazK2kZ7iMIc8o03YmHWKA7V5Q9Xt+5ffzcM7sNb7OJMXmYAZq74ds758bqdsZEFVJ5Bxf8wdBlRXItcsPml7z5Hm/PvfiMzhJ3NFivpAsifN7rxt4sDMKZd8ZxU32/8NKUKf+Q8M3r6WNSEdxAR4ejIXEHYpXI9iosRLOK63BJ9seewBKxd1rX8eKXwvd+UgK77UqIX+fgX+wOdEMf14phKggZtgooqiq9OD5sxkCKw1G9oQCuy+mQfaTQGaOq1RXwt7eR+IAw2RXdRdR6B6Nbv1iYFjs2RJKWdyilx3jI4iXpYrOpXcntEZ0XjUaOCalvpdLu5ITHWWkJbuaZ/3mz8+NRHJ2uuLoufYL1qYAXsTSSeF6bf5mSnKsT2b33R+YFZ2TGD3IL53UbsACdI3TFz19ewOfVpZxbdme21yaAEOCxtkahO9wU911PhEsLPSidaJiChC7HKuXTIFCnt8RukGxizQ/azbNw02f4SBK221giZR1zDArZGlb9iLSaq/uN9OvD9c7lzxdotM4sk1qExbOCYq9R1hqONKQ3wz330UDK32f7Xsww1jkGKPedvr2s+JdotbPFBpse2TfVdoOwuGQ/wPWlr3jwZP58jcNhISomuVV/mw1EgWb7JjwNNTtN/Auend/50EZdvs2EtXyrwHNYhuUajmfVorfkL/FdlVixa9EEZtZrDzBTv9L0y2XfZxu5SeO/GmnNpzoMPVFblCPAbZgvYsoZsfrM+yjlmSdLAsUnUNFQY+fL5UCPnVXG/EicwOBMp3LMs5m0janHggfAVcScYEIgNxbqs1YgYYU8URa3nn6xpe6IfOWH+yw4S88W/zHfkzOuEYDtW1RSLmrqRcSlPFO4Eii+52UHqB8CNGIkvFXjI0uhnm3fB8WFACIx+za2CTZkMCSF/3KkEOKtt6+UPPqV8jk4DkG/aCqnfio8Vu"
	data1, _ := base64.StdEncoding.DecodeString(data)
	// tool := NewAesTool(key, 16)
	r, err := ECBDecrypt(data1, key)
	// r, err := tool.Decrypt(data1)
	t.Log(string(r), err)
}
