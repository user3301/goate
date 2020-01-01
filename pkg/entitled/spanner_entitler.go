package entitled
import(
	"cloud.google.com/go/spanner"
)

type SpannerEntitler struct {
	client *spanner.Client
}

func NewSpannerEntitler(client *spanner.Client) (s *SpannerEntitler, err error) {
	s = &SpannerEntitler{client:client}
	return
}

