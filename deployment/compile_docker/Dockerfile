FROM amazonlinux

WORKDIR /

RUN yum update -y

RUN yum install wget tar gzip make -y

RUN wget -c https://golang.org/dl/go1.15.5.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.15.5.linux-amd64.tar.gz

RUN export PATH=$PATH:/usr/local/go/bin

RUN yum install golang -y

RUN go version

RUN go get -u -d gocv.io/x/gocv

RUN yum install unzip sudo -y

RUN yum install gcc-c++ -y

RUN g++ --version

RUN wget https://cmake.org/files/v3.6/cmake-3.6.2.tar.gz
RUN tar -zxvf cmake-3.6.2.tar.gz

WORKDIR /cmake-3.6.2

RUN sudo ./bootstrap

RUN sudo make

RUN sudo make install

RUN cmake --version

WORKDIR /

RUN yum install git -y

RUN git clone https://github.com/opencv/opencv.git

RUN mkdir opencv_build

RUN yum install pkgconfig -y

WORKDIR /opencv_build

RUN cmake /opencv -DOPENCV_GENERATE_PKGCONFIG=ON
RUN make
RUN sudo make install

RUN mkdir /usr/local/lib/pkgconfig
RUN export PKG_CONFIG_PATH=/usr/local/lib/pkgconfig

WORKDIR /usr/local/lib/pkgconfig

COPY opencv4.pc ./
RUN sudo chmod 777 opencv4.pc

WORKDIR /root/go/src/gocv.io/x/gocv

RUN ls -al /usr/local/lib/pkgconfig

# RUN make install
