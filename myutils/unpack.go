package myutils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

const MsgHeader = "12345678"

// 使用 binary.Write 写入数据
func Encode1(conn io.Writer, content string) (err error) {

	// 需要注意的是 MsgHeader 要转为 []byte， 不转的话语法不会报错,但是 server 端得不到正确的数据
	if err = binary.Write(conn, binary.BigEndian, []byte(MsgHeader)); err != nil {
		err = fmt.Errorf("write MsgHeader err:%v\n", err)
		return err
	}

	length := int32(len([]byte(content)))

	if err = binary.Write(conn, binary.BigEndian, length); err != nil {
		err = fmt.Errorf("write length err:%v\n", err)
		return err
	}

	// 需要注意的是 content 要转为 []byte， 不转的话语法不会报错,但是 server 端得不到正确的数据
	if err = binary.Write(conn, binary.BigEndian, []byte(content)); err != nil {
		err = fmt.Errorf("write content err:%v\n", err)
		return err
	}

	return nil
}

// 使用 conn.Write 写入
func Encode2(conn net.Conn, content string) (err error) {

	if _, err = conn.Write([]byte(MsgHeader)); err != nil {
		err = fmt.Errorf("write MsgHeader err:%v\n", err)
		return err
	}

	// 将 int 类型的长度转为 []byte 才能使用 conn.Write 写入
	length := int32(len([]byte(content)))
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, length)
	lengthBytes := bytesBuffer.Bytes() // 得到对应的 []byte 数据
	if _, err = conn.Write(lengthBytes); err != nil {
		err = fmt.Errorf("write lengthBytes err:%v\n", err)
		return err
	}

	if _, err = conn.Write([]byte(content)); err != nil {
		err = fmt.Errorf("write content err:%v\n", err)
		return err
	}

	return nil
}

// 使用 binary.Read 读取数据
func Decode1(conn io.Reader) (bodyBuf []byte, err error) {

	msgHeaderBuf := make([]byte, len(MsgHeader))
	if err = binary.Read(conn, binary.BigEndian, msgHeaderBuf); err != nil {
		err = fmt.Errorf("Decode1 MsgHeader err:%v\n", err)
		return nil, err
	}

	if string(msgHeaderBuf) != MsgHeader {

		err = fmt.Errorf("Decode1 get MsgHeader error:%v\n", err)
		return nil, err
	}

	lengthBuf := make([]byte, 4)
	if err = binary.Read(conn, binary.BigEndian, lengthBuf); err != nil {
		err = fmt.Errorf("Decode1 lengthBuf err:%v\n", err)
		return nil, err
	}

	length := binary.BigEndian.Uint32(lengthBuf)

	bodyBuf = make([]byte, length)
	if err = binary.Read(conn, binary.BigEndian, bodyBuf); err != nil {
		err = fmt.Errorf("Decode1 bodyBuf err:%v\n", err)
		return nil, err
	}

	return bodyBuf, nil

}

// 使用 io.ReadFull 读取数据
func Decode2(conn io.Reader) (bodyBuf []byte, err error) {
	msgHeaderBuf := make([]byte, len(MsgHeader))

	if _, err = io.ReadFull(conn, msgHeaderBuf); err != nil {
		err = fmt.Errorf("Decode2 msgHeaderBuf err:%v\n", err)
		return nil, err
	}

	if string(msgHeaderBuf) != MsgHeader {
		// 发送的数据开头不是 MsgHeader 的话马上就显示这个报错信息了
		err = fmt.Errorf(" Decode2 get MsgHeader error,%v", err)
		return nil, err
	}

	lengthBuf := make([]byte, 4)
	if _, err = io.ReadFull(conn, lengthBuf); err != nil {
		err = fmt.Errorf("Decode2 lengthBuf err:%v\n", err)
		return nil, err
	}

	length := binary.BigEndian.Uint32(lengthBuf)

	bodyBuf = make([]byte, length)
	if _, err = io.ReadFull(conn, bodyBuf); err != nil {
		err = fmt.Errorf("Decode2 bodyBuf err:%v\n", err)
		return nil, err
	}

	return bodyBuf, nil

}

// 使用 conn.Read 读取数据
func Decode3(conn net.Conn) (bodyBuf []byte, err error) {
	msgHeaderBuf := make([]byte, len(MsgHeader))

	if _, err = conn.Read(msgHeaderBuf); err != nil {
		err = fmt.Errorf("Decode3 msgHeaderBuf err:%v\n", err)
		return nil, err
	}

	if string(msgHeaderBuf) != MsgHeader {

		err = fmt.Errorf("Decode3 get MsgHeader error,%v", err)
		return nil, err
	}

	lengthBuf := make([]byte, 4)
	if _, err = conn.Read(lengthBuf); err != nil {
		err = fmt.Errorf("Decode3 lengthBuf err:%v\n", err)
		return nil, err
	}

	length := binary.BigEndian.Uint32(lengthBuf)

	bodyBuf = make([]byte, length)
	if _, err = conn.Read(bodyBuf); err != nil {
		err = fmt.Errorf("Decode3 bodyBuf err:%v\n", err)
		return nil, err
	}

	return bodyBuf, nil

}
