		return nil
		ctxLines := c.ctxLines
		if ctxLines > len(c.afterContext) {
			ctxLines = len(c.afterContext)
		}
		c.current.AddOp(Equal, c.afterContext[:ctxLines]...)
		c.beforeContext = c.afterContext[ctxLines:]