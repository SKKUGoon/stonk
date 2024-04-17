package kis

import "fmt"

func (c *KISClient) UsePrefixFn(fn AcntHandlerFunc) {
	c.preHandlers = append(c.preHandlers, fn)
}

func (c *KISClient) UseClosingFn(fn AcntHandlerFunc) {
	c.closingHandlers = append(c.closingHandlers, fn)
}

func (c *KISClient) SetTx(fn AcntHandlerFunc) {
	c.handlers = append(c.handlers, fn)
}

func (c *KISClient) Exec() (map[string]interface{}, error) {
	// Execute all prefix functions - inside the queue
	// Prefix functions are made with `UsePrefixFn`
	for i, pf := range c.preHandlers {
		_, err := pf()
		if err != nil {
			fmt.Printf("err during prefix handler %v: %v\n", i, err)
			return nil, err
		}
	}

	// Execute all main functions inside queue
	payload := map[string]interface{}{}
	for i, f := range c.handlers {
		if data, err := f(); err != nil {
			fmt.Printf("err during handler %v: %v\n", i, err)
			return nil, err
		} else if data != nil {
			// Some data (any or interface{}) was returned from executing function
			payload[fmt.Sprintf("payload%v", i)] = data
		}
	}

	// Re-initialize handlers
	c.handlers = []AcntHandlerFunc{}
	return payload, nil
}
