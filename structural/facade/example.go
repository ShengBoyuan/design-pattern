package facade

import "fmt"

// Complex Subsystem
type VideoFile struct{}
type MPEG4CompressionCodec struct{}
type OggCompressionCodec struct{}
type CodecFactory struct{}
type BitrateReader struct{}
type AudioMixer struct{}

func (CodecFactory) Extract(File) []byte {
	return nil
}

func (MPEG4CompressionCodec) Handle(data []byte) []byte {
	return data
}

func (OggCompressionCodec) Handle(data []byte) []byte {
	return data
}

// Facade
type VideoConverter struct{}
type File struct {
	Name string
	Data []byte
}

func (VideoConverter) Convert(fileName string, format string) File {
	file := File{Name: fileName}
	codecFactory := CodecFactory{}

	sourceCodec := codecFactory.Extract(file)
	if format == "mp4" {
		file.Data = (&MPEG4CompressionCodec{}).Handle(sourceCodec)
	} else {
		file.Data = (&OggCompressionCodec{}).Handle(sourceCodec)
	}

	// Omit some logical combination of subsystem classes.
	return file
}

func (f File) Save() {
	fmt.Println("file", f.Name, "has been saved!")
}

// Additional Facade

// Client
func ChangeFileFormat() {
	convertor := &VideoConverter{}
	mp4 := convertor.Convert("youtubevideo.ogg", "mp4")
	mp4.Save()
}
