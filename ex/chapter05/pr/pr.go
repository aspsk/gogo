package pr

func pr() (ret int) {

	defer func() {
		switch recover() {
			default:
				fallthrough
			case nil:
				panic("no way!")
			case "":
				ret = 0xdeadbeef
		}
	}()

	panic("")

}
