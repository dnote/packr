package resolver

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_HexGzip_Find(t *testing.T) {
	r := require.New(t)

	x, err := HexGzipString("foo!")
	r.NoError(err)
	files := map[string]string{
		"foo.txt": x,
	}
	d, err := NewHexGzip(files)
	r.NoError(err)

	f, err := d.Resolve("", "foo.txt")
	r.NoError(err)

	fi, err := f.FileInfo()
	r.NoError(err)
	r.Equal("foo.txt", fi.Name())

	b, err := ioutil.ReadAll(f)
	r.NoError(err)
	r.Equal("foo!", strings.TrimSpace(string(b)))
}

func Test_fo(t *testing.T) {
	r := require.New(t)
	x := "1f8b08000000000000ff9c575f6fe3b8117fd7a798ea490e6c0a7d2b5ce4806c76b3b9c29718916f53e0b04818696cb34b715892b2e306feeec5905262a7ebddb6010249e4f037bff9c399b195f537b942907550647c96a9d6920b506400f94a8575f7246a6acb153d75cba5d454a2d9ecf293bbfd930596e46af45e9f906cb7934182a5ad74b2d5b45aa13b7de040e83d85ce38d21a9bd263dd39ccb30ce0e505d41204d9e0c5bd0aeb3959d8eff9e0cb4b5afd8d1ad47efe6db5dfffd4a689255bf6ffed364ff8681a98ecf7efb55df87b7c8a1b00b577cbd326f12e63a93fffc59c96e2ddd3043988ee67842ee6bff6840e419c2f6b72fe34f673e9d17bce0d96a9c9043421ec2cfec0a437a177a446595696f0e9e60b280f9dc70602c11ab505bf55a15e83c7109459797892bc4b06b66b7408618d7c525aab552d395519e109955981eb8c808fb8949d0ebc9a37b8414db6451372916da48b1acf8133577cc650e49f6f1f3edd7cc9c7c7b2a3282bad85b3de1a71616df6fdb8b2e802ce383062e1a4f15a0672596f2c07a02ce1c25a66946c905a83a32ea007691a6855d368dc4a877c53a0d7c8a7fc9a3addc01342834b65b011b0582bcf48618d60d06d106a34011dd01276d4b977ce115959f2d21d75eccef181b231ac1c75d68f01432d840026d060ada5c30616b77398fc021f6fef6f1823e96d511acf3ed85107b269401e920f048fd2da473893cb80eeac07e3c8488688eac610d632a477d82aade1e676016bb9c1b461707b80c9f62278d9c6a0b3d9ae4336d5a1a7ced503e168ab07e9618b5a0f46ff1a406a4f3ded083fb8dd21d46bacbf6103ca4467926bd0f1dbeec811d1f8c78a3d7da534fa476621e1f26271793db998cd12e2183c810a43c0a4deca1d67251fb65ad6d880963e0cbae29923ea63e6ae42f28875b8413390658cd7b044cfb2602030b841c7c9514bae77225b76a6e64c2b4647790b2f5c5596319fcfcfc1281d57202d0cf9266e705b0cefb736b24a62009fcc66ca5767dc7f9faa27fc57a5225105723885a164889b4eeb97fd003077782f1dfa29fcf175d0d9af0d3ab9c4382ffaeb5c8cc4b5348d463740ec0fc80c55e588c18d6c710af9c350e1f95b700cf7fb879e569e30f6a32c3ef99e76815a1914bb74070e1be5b08edeaeaad9e034f1bbc722b6acaa9a15a3b7d3335a81c37f76e8436a6318d07928964ac717be98bb9138c239e864623e9c99c5ef1ef8b4b7cb122a0c29a57ab57dc985589803c1dfaadb9b237d073559545c020faa45f90f4f261f8db203bfa60e72baa94516734781fd245752191fe0b2babb021982acbf7901eb10ac9f96e576bb15b495de0a72ab5299069f855ddbf2d291f7934a057cb84b663c5c915ba1db3d148c341ad4dc614b9b6856a3bc7cd25c33943ff627b751cee59f1a71340744f87b27ad0794f5fad59dca8084c0553dcd456290855a7c91bac3220fcff948146796acb8246330cafd4f94e318917a47d252b4711a111f3ffc1fb1a83074363696ce63e2ae5391991e693ddc29beaf87453f7f5a1479998fe19a5aecefe07713f390061f7cab99e9bcf41e83ff40cf2366e963ef5af2362c1db531896df7a4550de9d291db65ff71bf233187a173869564fbec44572ecb23d35361d5241bd0544bddab1ef39c019d8ddac36bf386c77e0a168bc771eca8a6490883ea777d2f7a9a123e43d5e462dd1874918b69c5507d66c5bec2eb2d39046596e4da4875fa7a5d5e4729a178d82e1baa7d1911d5bfa2682af8c7717c2de7bfbd12bc62292eab3ca9a08bffe4524f588ce3ca799c3a630b882324bf7da0e722172269449f8fc690a399fc5ee5a3bfc6437f7ad74b4415c816e8dc2806ea354e8b032ec528dbbf1b8d8642fa030773e7ee77fb822c0d2853531bc7bee4d138212cb9372a0f86025c2f16f34a40cefe9c96253ecbd66ae4293587f35fd2ba7fb77130ebbce96f083d9c9d190a67678026dee2aa9a89183f1eba8e262e581034144933d8966b734d6d8bf146c66908aca3e7dd4f221d65fe8b2c39f9cba74cf9f1d6a87e941bbddf871f6287214b68c7334155cdeefa584cd388c083f539e4d651d3c5cbd337d7aa9acdd9926b940d3a3f8556da3f7ce0b1f06b7abce47f9f5c91db4ad76033e15642f9b40f4f1edbfc9eb3e6df010000ffffa0d4ca87a10e0000"
	y, err := UnHexGzipString(x)
	r.NoError(err)
	r.NotEqual("", y)
}
