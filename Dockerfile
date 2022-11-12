FROM ubuntu

WORKDIR /qiniu

COPY . /qiniu


RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal main restricted > /etc/apt/sources.list
RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-updates main restricted >> /etc/apt/sources.list
RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal universe >> /etc/apt/sources.list
RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-updates universe >> /etc/apt/sources.list
RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal multiverse >> /etc/apt/sources.list
RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-updates multiverse >> /etc/apt/sources.list
RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-backports main restricted universe multiverse >> /etc/apt/sources.list
RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-security main restricted >> /etc/apt/sources.list
RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-security universe >> /etc/apt/sources.list
RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-security multiverse >> /etc/apt/sources.list


RUN apt-get update -y


RUN apt install tzdata -y

CMD cd back_end/ && ./main

EXPOSE 8080
