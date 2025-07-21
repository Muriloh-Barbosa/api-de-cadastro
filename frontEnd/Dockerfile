#arquivos estáticos com Nginx

# imagem
FROM nginx:alpine

# Remover config padrão do Nginx
RUN rm -rf /usr/share/nginx/html/*

# Copia dos arquivos do frontend para o diretório público do Nginx
COPY . /usr/share/nginx/html

# Expõe a porta 80 para acesso HTTP
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]