import 'package:diox/diox.dart';
import 'package:retrofit/retrofit.dart';

import '../../domain/entities/torrent/torrent.dart';

part 'torrent_api_client.g.dart';

@RestApi()
abstract class TorrentAPIClient {
  factory TorrentAPIClient(Dio dio) = _TorrentAPIClient;

  @GET('/1337x')
  Future<List<Torrent>> get1337x(@Query('q') String query);

  @GET('/yts')
  Future<List<Torrent>> getYTS(@Query('q') String query);

  @GET('/rarbg')
  Future<List<Torrent>> getRarbg(@Query('q') String query);

  @GET('/lime')
  Future<List<Torrent>> getLime(@Query('q') String query);

  @GET('/kat')
  Future<List<Torrent>> getKat(@Query('q') String query);

  @GET('/tpb')
  Future<List<Torrent>> getTPB(@Query('q') String query);

  @GET('/tgx')
  Future<List<Torrent>> getTGX(@Query('q') String query);
}
