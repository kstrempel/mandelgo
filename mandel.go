package main

import ("image";
        "image/png";
        "image/color";
        "bufio";
        "fmt";
        "os";
        "math";
        "time"; )

func PointIteration(cx,cy,maxValue float64,maxIter uint8) uint8{

	quadValue := float64(0.0);
	iter := uint8(0);
	x := float64(0.0);
	y := float64(0.0);
	
	for quadValue <= maxValue && iter < maxIter {
		xt:= ( x * x ) - ( y * y) + cx;
		yt:= ( float64(2.0) * x * y ) + cy;
		x = xt;
		y = yt;
		iter++;
		quadValue = ( x * x ) + ( y * y );
	}

	return iter;
}

func main(){
	start := time.Now();     
	
	const pictureSize = 5000;
	img := image.NewRGBA(image.Rect(0,0,pictureSize,pictureSize));
	
	f, err := os.OpenFile("mandel.png", os.O_WRONLY|os.O_CREATE, 0666);

	if err != nil {
		fmt.Printf("Can't create picture file\n");
	}
	
	deltaX := math.Abs(float64(-2.0 - 1.0)) / float64(pictureSize);
	
	deltaY := math.Abs(float64(-1.0 - 1.0)) / float64(pictureSize);
	
	cx := float64(-2.0);
	for x:=0;x<pictureSize;x++ {
		cx+=deltaX;
		cy := float64(-1.0);
		for y:=0;y<pictureSize;y++ {
			cy+=deltaY;
			iter := PointIteration(cx,cy,255.0,255);
			
			col := color.RGBA{255,iter,iter,iter};
			col.A = 255;
			col.R = iter;
			col.G = iter;
			col.B = iter;
			img.Set(x,y,col);
		}
	}

	w := bufio.NewWriter(f);
	png.Encode(w,img);
	w.Flush();
	
	fmt.Printf("Seconds needed %d\n",time.Now().Sub(start));
}
