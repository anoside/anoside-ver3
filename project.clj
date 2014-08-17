(defproject anoside "0.1.0-SNAPSHOT"
  :description "匿名でつぶやきが投稿できるWebサービス"
  :url "http://anoside.com"
  :dependencies [[org.clojure/clojure "1.6.0"]
                 [org.clojure/clojurescript "0.0-2311"]]
  :plugins [[lein-cljsbuild "1.0.3"]]
  :cljsbuild {
    :builds [{
      :source-paths ["web/src"]
      :compiler {
        :output-to "web/public/javascripts/hello.js"
        :optimizations :whitespace
        :pretty-print true}}]})
