# Artie (Copy Service)

Artie is a project whose goal is to create an application suite for
orchestrating the copying, transcoding, and archiving of DVDs and Blu-rays for
so that they can be streamed using a service like Jellyfin.

This go module provides a service for copying. Each service instance controls a
single DVD/Blu-ray drive. Under the hood, it uses MakeMKV to create one or more
MKV files and handed off to the next stage in the process.

