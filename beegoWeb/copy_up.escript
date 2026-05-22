#!/usr/bin/env escript
%% -*- erlang -*- 
%%! -smp enable -pa ../ebin
%% escript version_up -id 2 -serverip "127.0.0.1" -serverdir "/home/game/server_version/sever/game_server/" -serververion 1

main(Options)->
	code:add_path("./ebin/"),
	io:format("main: ~p ~n", [Options]),
	OptionList = get_arguments(Options),
	io:format("version_up option: ~p ~n", [OptionList]),
	
	{ok, Dir} = file:get_cwd(),
	VersionFile = [Dir, "/server_version/version.txt"],	
	
	{"id", [ServerId]} = lists:keyfind("id", 1, OptionList),
	{"serverip", [ServerIp]} = lists:keyfind("serverip", 1, OptionList),
	{"serverdir", [ServerDir]} = lists:keyfind("serverdir", 1, OptionList),
	{"serververion", [ServerVersion]} = lists:keyfind("serververion", 1, OptionList),
	VersionList = case file:consult("./server_version/version.txt") of
					{ok, [Terms]}->
						Terms;
					{error, Reason}->
						io:format("file:consult ############## ~p, ~s ~n", [Reason, VersionFile]),
						[]
				end,
	filelib:ensure_dir(ServerDir),
	VersionDir = [Dir, "/server_version/", ServerVersion, "/*"],
	if 
		ServerIp =:= "127.0.0.1"->
			os_util:run_exe("unalias cp"),
			Cp = ["cp -rf ", VersionDir, " ", ServerDir];
			%%UpdateShell = [ServerDir, "/script/", "version_update.sh"];
		true->
			Cp = ["scp -r -P 29999 ", VersionDir," root@", ServerIp, ":",ServerDir]
			%%Update = [ServerDir, "/script/", "version_update.sh"],
			%%UpdateShell = ["ssh -p 29999 root@", ServerIp, " \" ", Update, "\" "]
		end,
	_Result = os_util:run_exe(Cp),
	%%Result1 = os_util:run_exe(UpdateShell),
	IntServerId = list_to_integer(ServerId),
	AtomServersion = list_to_atom(ServerVersion),
	case lists:keyfind(IntServerId, 1, VersionList) of
		false->
			NewVersionList = [{IntServerId, AtomServersion}|VersionList];
		{IntServerId, _}->
			NewVersionList = lists:keyreplace(IntServerId, 1, VersionList, {IntServerId, AtomServersion})
	end,
	MapString = lists:map(fun({TmpServerId, TmpServerVersion})->
						"{" ++ integer_to_list(TmpServerId) ++ "," ++ transform(TmpServerVersion) ++ "}"
					end, NewVersionList),
	FileContent = "[" ++ string:join(MapString, ",") ++ "].",
	file:write_file(VersionFile, unicode:characters_to_binary(FileContent)).

transform(Sv) when erlang:is_atom(Sv) ->
	atom_to_list(Sv);
transform(Sv) when erlang:is_integer(Sv) ->
	integer_to_list(Sv);
transform(Sv)->
	util_string:term_to_string(Sv).
	
get_arguments(Options)->
	Fun = fun(Opt,{ThisKV,KVs})->
				  case Opt of
					  "-"++ Opt1->				
						  NewKVs = case ThisKV of
									   []-> KVs;
									   _-> KVs ++ [ThisKV]
								   end,
						  NewThisKV = {Opt1,[]},
						  {NewThisKV,NewKVs};
					  _->
						  NewKVs = KVs,
						  case ThisKV of
							  []->
								  {[],NewKVs ++ [{Opt}]};
							  {LastKey}-> 
								  NewThisKV = {LastKey,[Opt]},
								  {NewThisKV,NewKVs};
							  {LastKey,LastVal}->
								  NewThisKV = {LastKey,LastVal++[Opt]},
								  {NewThisKV,NewKVs}
						  end
				  end
		  end,
	{ThisKV,KVs} = lists:foldl(Fun, {[],[]}, Options),
	KVs ++ [ThisKV].
