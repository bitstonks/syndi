package generators

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bitstonks/syndi/internal/config"
)

// yaay, globals!
var lipsum = strings.Replace(`
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi odio nisi, iaculis at auctor et, venenatis id nisl. In ultrices enim eu ultricies facilisis. Curabitur egestas turpis mi. Ut eu scelerisque ipsum, eu gravida sapien. Pellentesque ac odio et orci congue condimentum blandit vel ex. Nunc vulputate feugiat ornare. Maecenas non gravida urna. Nullam vitae venenatis libero.
Aenean blandit semper quam, ac dictum dui molestie a. Proin ac purus imperdiet, auctor est et, ultricies arcu. Suspendisse lectus nisi, pellentesque eu rhoncus eu, placerat vitae sapien. Integer elementum dui sit amet elit auctor mollis. Donec consectetur ipsum lorem, ut efficitur felis elementum vel. Nulla facilisi. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; In ultrices posuere odio eu sollicitudin. Interdum et malesuada fames ac ante ipsum primis in faucibus. Donec aliquam sapien vel nisi sagittis, commodo accumsan turpis sodales.
In sed lacinia enim, eget malesuada urna. Praesent id ex eu purus venenatis sagittis. Donec auctor magna et nibh mollis, sed sodales tortor dignissim. Pellentesque facilisis placerat ante et porta. Morbi nec molestie nulla, et sollicitudin diam. Nam ut nisi non dolor fermentum posuere eget sit amet arcu. Integer condimentum dictum erat, nec pulvinar mauris pellentesque vel. Vestibulum mollis in turpis sit amet pharetra. Etiam id mattis velit. Donec sodales mauris vel nibh pellentesque faucibus. Curabitur quis euismod lorem. Nam suscipit ultrices erat, sed tristique urna porta quis. Vivamus finibus sit amet ipsum non pharetra.
Nullam ut aliquam est, et lobortis urna. Vestibulum tellus est, congue at metus sed, consequat mollis est. Nullam sapien sem, porttitor in mollis ut, porta at justo. Vestibulum tincidunt eu lectus volutpat gravida. Sed ultrices iaculis neque non facilisis. Morbi malesuada lacus at sapien elementum viverra. Vestibulum vel quam ac turpis sollicitudin dictum vitae eu libero. Proin in mi dictum, varius lorem sed, vehicula magna. Vivamus dolor nisl, posuere et erat at, tincidunt feugiat ex. Praesent ornare eros nibh, non bibendum nibh fermentum consectetur. Donec id mauris euismod, efficitur turpis non, tristique ex. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Etiam imperdiet neque et tortor accumsan, non tincidunt ex dapibus.
Fusce facilisis ligula in tellus molestie, quis aliquam metus dignissim. Integer odio diam, condimentum eget neque sit amet, laoreet volutpat mauris. Nunc a auctor leo. Nulla nec fringilla massa. Vestibulum laoreet lacus quis lorem imperdiet laoreet. Cras ornare neque tellus, vitae gravida nisl ultricies vel. Phasellus consequat tortor dolor, sed varius velit vestibulum in. Fusce lacinia, mi vel viverra viverra, lacus velit vulputate justo, nec vehicula ipsum enim et ligula. Sed sed convallis ex. Nam lobortis arcu id magna tincidunt, nec interdum mi ultrices. Phasellus et turpis sed arcu sodales pretium eu non magna. Ut dapibus leo tincidunt, vulputate nunc sed, interdum lacus.
Quisque erat sem, ultrices et orci non, tincidunt iaculis tortor. Praesent sit amet enim id massa iaculis malesuada vitae sit amet enim. Quisque sagittis placerat elit semper porttitor. Aliquam egestas, erat vitae dignissim luctus, enim erat interdum libero, in blandit nibh dui nec urna. Integer imperdiet malesuada sapien, non facilisis felis fermentum vel. Vivamus sit amet purus pulvinar, finibus lectus a, tempor nisl. Suspendisse tempor ultricies facilisis. Sed ut felis magna.
In scelerisque sollicitudin auctor. Quisque eu lorem nec nibh porta faucibus ac sit amet augue. Mauris imperdiet massa vel sem gravida, eget lobortis orci ultricies. Morbi mattis leo purus, sed consequat tellus pharetra efficitur. Proin felis augue, mattis eget bibendum et, tempus ut nulla. Sed a hendrerit nisl. Sed nulla orci, porttitor sed sapien vel, hendrerit dictum felis. Integer diam purus, tincidunt eget aliquet id, interdum eu nisl. Aliquam posuere est nunc, et pulvinar nulla fermentum sollicitudin. Etiam interdum tellus ut orci venenatis, et ultrices tellus accumsan. Sed suscipit sagittis leo, a imperdiet tellus. Sed in commodo tortor, eget sollicitudin felis. Ut accumsan gravida mi, sed venenatis mauris tincidunt quis. Sed ut mattis nulla.
Maecenas turpis lectus, ullamcorper sit amet arcu in, sollicitudin venenatis ex. Proin finibus tortor nec fringilla tincidunt. Donec et nibh sit amet purus dictum malesuada. Quisque placerat viverra efficitur. Quisque quis molestie dolor, eu posuere augue. Praesent vitae imperdiet metus, id pellentesque sem. Praesent lacinia lobortis interdum. Fusce feugiat ligula at rhoncus fermentum. Curabitur vel accumsan urna. Vestibulum aliquam tortor lorem, eget ullamcorper arcu euismod et. Sed lectus nulla, elementum nec consectetur non, ornare nec arcu. Curabitur ac purus ac arcu mattis pulvinar. Aliquam facilisis lacus vitae mauris fringilla mollis.
Nulla facilisi. Fusce mattis, mi sit amet placerat convallis, dui nisi feugiat mi, et euismod mi dui vel velit. Mauris egestas dictum quam nec pellentesque. Vivamus quis fermentum augue. Donec lobortis hendrerit libero, eu vestibulum dolor malesuada nec. Sed sit amet ornare mauris, at varius sem. Ut sed aliquam sapien. Fusce in enim iaculis, laoreet est ac, tincidunt lorem. Ut sit amet leo a tortor finibus faucibus. Aenean bibendum mauris ut scelerisque vulputate. Nunc nunc orci, congue vitae metus maximus, hendrerit semper ante. Aenean convallis felis at erat convallis ultricies. Nulla luctus, magna vel volutpat tempor, felis felis convallis tortor, at finibus turpis tortor at mauris. Vestibulum in eros vel erat dapibus auctor. Nam tellus sem, auctor in cursus eget, facilisis in ex.
Cras gravida orci vel nisl tincidunt iaculis. Nulla fringilla ultricies est, ac vulputate turpis tempor sit amet. Vestibulum viverra malesuada nisi ut sollicitudin. Suspendisse orci urna, sagittis bibendum eleifend eget, hendrerit sed enim. Integer sed dui euismod sapien iaculis rhoncus et sollicitudin sapien. Maecenas quam turpis, finibus at est non, eleifend semper tortor. Proin ornare eros at augue fringilla venenatis. Cras commodo enim non lacus scelerisque, vitae condimentum nisl sagittis. Aliquam pretium ipsum vitae tincidunt volutpat. Nulla aliquam lobortis nulla placerat aliquam.
Aliquam ullamcorper felis ex, eget convallis felis pretium et. Mauris tempor ac quam in viverra. Fusce non pretium arcu. Duis mattis quam ac leo accumsan, ut faucibus ipsum elementum. Proin at tellus non odio dictum ullamcorper non in arcu. Sed a pellentesque dui. Donec elit neque, accumsan quis nisl ac, consequat convallis sem. In iaculis leo tortor, eu dictum ipsum commodo non.
Sed bibendum massa ut orci aliquet tincidunt. Nam vulputate metus vitae tempor vestibulum. Vivamus ut venenatis arcu. Cras ut tristique mauris, non luctus odio. Nullam et urna consectetur, vulputate ipsum sed, aliquam metus. Quisque ac lectus libero. Nullam accumsan nisl at urna faucibus, non scelerisque lacus ultricies. Sed nunc sapien, porttitor venenatis sollicitudin ac, bibendum et nulla. Aliquam hendrerit vulputate lectus, a vulputate orci mattis eget. Vivamus bibendum imperdiet urna ut luctus. Sed ac sapien ullamcorper, scelerisque sem rhoncus, scelerisque eros. Morbi vel tortor ullamcorper, sollicitudin velit in, pulvinar tortor. Interdum et malesuada fames ac ante ipsum primis in faucibus. Vestibulum et ex lorem. Pellentesque tincidunt nec elit vel dapibus.
Curabitur lobortis lacus nec ante venenatis ornare in quis tellus. In nec fermentum metus, sed viverra lectus. Nam dapibus augue in leo mattis, id lacinia dolor fringilla. Phasellus facilisis nisl ut ullamcorper lacinia. Phasellus eleifend vitae tortor ut ornare. Donec vel sapien auctor, pharetra neque nec, sagittis est. Sed mattis augue interdum arcu convallis, sodales viverra tortor vehicula. Donec luctus nisi commodo massa bibendum, at tincidunt elit lacinia. Phasellus a turpis pulvinar, sodales massa eu, rutrum eros. Donec pharetra nulla dui, id rutrum lacus placerat non. Nulla luctus enim dui, non pellentesque ex auctor nec. Fusce dignissim a ipsum vitae lacinia. Aenean vel scelerisque nisl, eu hendrerit ligula. Duis lacinia, ipsum sit amet pulvinar posuere, nisi nisl egestas eros, vitae maximus elit massa sit amet leo. Donec pharetra nec lorem vel finibus.
Praesent facilisis auctor convallis. Ut accumsan tortor et ipsum eleifend malesuada. Aliquam eget tortor euismod, vestibulum diam ac, euismod turpis. Suspendisse dapibus, nisi nec lobortis pretium, ligula nisi scelerisque nisi, quis dictum nulla tellus sit amet purus. Nullam ac enim eget odio bibendum pulvinar. Proin a dui vitae sem tincidunt pellentesque nec sed justo. Aliquam eget magna nunc. Morbi ullamcorper nunc nec sem mollis bibendum. Aliquam facilisis justo erat, id ultricies nunc molestie sit amet. Sed a porta magna, sed finibus risus. Morbi justo dui, faucibus ut nisl et, placerat gravida turpis. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac tortor in eros tristique vehicula vitae ut nunc.
Duis ultrices augue non urna ornare sollicitudin. Ut vel ante metus. Nam imperdiet, enim ac ornare rutrum, nibh tellus facilisis velit, ut ultrices massa libero a magna. Cras feugiat egestas nunc, nec maximus diam. Ut et mi enim. Morbi posuere elementum volutpat. Aenean quis tortor quis elit consequat dictum. Vestibulum ornare rhoncus finibus. Nam ac lorem convallis, lacinia dui non, vulputate nunc. Pellentesque odio nisi, volutpat quis auctor sed, fringilla eget mi.
Fusce eleifend tellus velit, eget mattis ipsum consectetur sit amet. Phasellus ante sapien, dignissim eget augue sed, condimentum pellentesque erat. Etiam sed ante dolor. Nulla ut enim mollis, sagittis arcu ac, egestas leo. Fusce fermentum, ex et molestie pellentesque, nunc dui lobortis quam, et lobortis dui nisl in nisi. Aliquam in nisl placerat, tincidunt nisi ac, aliquet felis. Nunc congue urna at est congue iaculis. Suspendisse nec laoreet eros. Nullam at luctus sapien. Nullam sem purus, malesuada quis aliquam eget, pulvinar et nisl. Mauris velit purus, fringilla vitae volutpat vitae, fringilla sit amet velit.
Nullam sed ipsum vestibulum, hendrerit diam vel, malesuada augue. Aenean luctus venenatis leo, in vestibulum tortor fringilla volutpat. Sed fringilla, tellus vel ullamcorper consectetur, est nunc auctor enim, quis sodales turpis est in enim. Morbi gravida facilisis leo, vitae cursus urna. Sed lobortis sit amet urna et maximus. Fusce dignissim, enim et elementum dapibus, nulla mi auctor arcu, a tincidunt sem elit pulvinar quam. Praesent sodales hendrerit interdum.
Donec vitae fermentum metus. Sed dictum leo vel ipsum viverra consectetur. In hac habitasse platea dictumst. Donec tristique dolor a elit convallis, faucibus volutpat lacus dapibus. Maecenas porta cursus posuere. Sed vitae purus a ex vestibulum pellentesque. Nulla vestibulum, metus id faucibus aliquet, dolor dolor egestas arcu, consequat condimentum nunc purus in risus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Fusce lobortis molestie pretium. Nunc elit justo, tristique vitae neque at, varius fermentum augue. Curabitur aliquam lacus erat, at semper odio consectetur pharetra. Nullam porta tempor eros nec vulputate. Quisque semper, sem nec commodo tincidunt, elit urna pellentesque nulla, ut volutpat mauris urna a lorem.
Mauris pellentesque pulvinar pellentesque. Aenean nec tortor et mauris hendrerit iaculis ac at dolor. Integer varius rhoncus semper. Aliquam erat volutpat. Aenean varius arcu eu sem pharetra, ut vulputate neque fringilla. Vivamus a convallis est. Donec molestie ut nisi vitae iaculis. Ut sagittis leo sit amet purus sagittis cursus. Sed porta erat lorem, sed ullamcorper ex lobortis vel. Pellentesque cursus pellentesque lorem, sit amet malesuada quam consequat quis. Sed tristique tristique leo. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean urna quam, porta sit amet auctor pharetra, accumsan eu ex. Fusce aliquam accumsan justo, a suscipit lacus scelerisque non. Aliquam erat volutpat.
Suspendisse ac dui lacus. Integer lectus nisi, congue ac mi quis, porta lobortis arcu. Donec aliquam velit luctus sagittis blandit. Donec commodo nisi vel aliquet placerat. Nulla facilisis tellus vitae nulla viverra cursus. Donec condimentum commodo dolor, vel dictum lacus viverra quis. Nam non lorem sed quam ultrices egestas ultrices sed ipsum. Suspendisse potenti.
`, "\r\n", " ", -1)
var lipsumLen = len(lipsum)

// TODO: what if Len is actually greater than lipsumLen?
type textGenerator struct {
	rng *rand.Rand
	len int
}

func NewTextGenerator(args config.ColumnDef) Generator {
	return &textGenerator{
		rng: newRng(),
		len: args.Length,
	}
}

func (g *textGenerator) Next() string {
	i := g.rng.Intn(lipsumLen - g.len)
	return fmt.Sprintf("'%s'", lipsum[i:i+g.len])
}
