package escpos

import "github.com/alexbrainman/printer"

type WindowsPrinter struct {
	ptr *printer.Printer
}

func NewWindowsPrinterByPath(printerPath string) (*Printer, error)  {
	p, err := printer.Open(printerPath)
	var wp WindowsPrinter
	if err != nil {
		return &Printer{
			s: wp,
			f: nil,
		}, err
	}
	wp.ptr = p
	wp.ptr.StartRawDocument("ticket.txt")

	return &Printer{
		s: wp,
		f: nil,
	}, nil
}

func (wprinter WindowsPrinter) Write(p []byte) (n int, err error) {
	return wprinter.ptr.Write(p)
}

func (wprinter WindowsPrinter) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (wprinter WindowsPrinter)  Close() error {
	wprinter.ptr.EndDocument()
		return wprinter.ptr.Close()
}