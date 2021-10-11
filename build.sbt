val AkkaVersion = "2.6.8"
val AkkaHttpVersion = "10.2.6"
run / fork := true
libraryDependencies ++= Seq(
  "com.typesafe.akka" %% "akka-actor-typed" % AkkaVersion,
  "com.typesafe.akka" %% "akka-stream" % AkkaVersion,
  "com.typesafe.akka" %% "akka-http" % AkkaHttpVersion,
  "ch.qos.logback"                 % "logback-classic" % "1.2.6"
)
