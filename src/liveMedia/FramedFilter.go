package liveMedia

type FramedFilter struct {
	FramedSource
	inputSource IFramedSource
}

func (this *FramedFilter) InitFramedFilter(inputSource IFramedSource) {
    this.inputSource = inputSource
}
